package mutationparser

import (
	"hcc/piccolo/action/graphql/queryparser"
	"hcc/piccolo/action/grpc/errconv"
	"hcc/piccolo/dao"
	"hcc/piccolo/lib/logger"
	"hcc/piccolo/lib/mysql"
	"hcc/piccolo/model"
	"strings"

	uuid "github.com/nu7hatch/gouuid"
	"innogrid.com/hcloud-classic/hcc_errors"
)

// SignUp : Do user sign up process
func SignUp(args map[string]interface{}) (interface{}, error) {
	groupID, groupIDOk := args["group_id"].(int)
	id, idOk := args["id"].(string)
	password, passwordOk := args["password"].(string)
	name, nameOk := args["name"].(string)
	email, emailOk := args["email"].(string)

	if !groupIDOk || !idOk || !passwordOk || !nameOk || !emailOk {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need id and group_id, password, name, email arguments")}, nil
	}

	if strings.ToLower(id) == "master" {
		logger.Logger.Println("SignUp(): Someone tried to sign up with master ID.")
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Hey, you can't be the master!")}, nil
	}

	sql := "select id from user where id = ?"
	row := mysql.Db.QueryRow(sql, id)
	err := mysql.QueryRowScan(row, &id)
	if err == nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "Provided ID is in use")}, nil
	}

	_, err = dao.ReadGroup(groupID)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, "Provided Group ID is not exist")}, nil
		}

		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	out, err := uuid.NewV4()
	if err != nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloInternalUUIDGenerationError, err.Error())}, nil
	}
	UUID := out.String()

	user := model.User{
		UUID:    UUID,
		GroupID: int64(groupID),
		ID:      id,
		Name:    name,
		Email:   email,
	}

	sql = "insert into user(uuid, group_id, id, password, name, email, login_at, created_at) values (?, ?, ?, ?, ?, ?, now(), now())"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLPrepareError, err.Error())}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err = stmt.Exec(user.UUID, user.GroupID, user.ID, password, user.Name, user.Email)
	if err != nil {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, err.Error())}, nil
	}

	user.Errors = errconv.ReturnHccEmptyErrorPiccolo()

	return &user, nil
}

// Unregister : Do user unregister process
func Unregister(args map[string]interface{}) (interface{}, error) {
	id, idOk := args["id"].(string)

	if !idOk {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "need a id argument")}, nil
	}

	if strings.ToLower(id) == "admin" || strings.ToLower(id) == "administrator" {
		logger.Logger.Println("Unregister(): Someone tried to unregister one of administrative ID.")
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloGraphQLArgumentError, "You can't delete administrative IDs")}, nil
	}

	user, _ := queryparser.User(args)
	if len(user.(model.User).Errors) != 0 && user.(model.User).Errors[0].ErrCode != 0 {
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, "user not found")}, nil
	}

	sql := "delete from user where id = ?"
	stmt, err := mysql.Prepare(sql)
	if err != nil {
		errStr := "Unregister(): " + err.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, errStr)}, nil
	}
	defer func() {
		_ = stmt.Close()
	}()
	_, err2 := stmt.Exec(id)
	if err2 != nil {
		errStr := "Unregister(): " + err2.Error()
		logger.Logger.Println(errStr)
		return model.User{Errors: errconv.ReturnHccErrorPiccolo(hcc_errors.PiccoloMySQLExecuteError, errStr)}, nil
	}

	return model.User{ID: id, Errors: errconv.ReturnHccEmptyErrorPiccolo()}, nil
}
