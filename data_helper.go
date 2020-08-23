package main

import (
	"errors"
	"strconv"
)

func getUserByEmail(n string) (user, error) {
	db, err := getDbConnection()
	if err != nil {
		return user{}, errors.New("could not get database connection")
	}
	defer db.Close()
	row := db.QueryRow("SELECT Id, Displayname, Email, Password, Locked, Admin FROM user WHERE Email = ?", n)
	var u user
	//var Locked int
	//var Admin int
	err = row.Scan(&u.Id, &u.Displayname, &u.Email, &u.Password, &u.Locked, &u.Admin)
	if err != nil {
		return user{}, errors.New("could not scan")
	}

	return u, nil
}

func getBuildDefCaption(id int) string {
	db, err := getDbConnection()
	if err != nil {
		return "could not fetch"
	}
	defer db.Close()
	var name string
	row := db.QueryRow("SELECT caption FROM build_definition WHERE id = ?", id)
	err = row.Scan(&name)
	if err != nil {
		return "could not scan"
	}
	return name
}

func getUserById(id int) (user, error) {
	var u user
	db, err := getDbConnection()
	if err != nil {
		return u, err
	}
	defer db.Close()
	row := db.QueryRow("SELECT Id, Displayname, Email, Admin FROM user WHERE Id = ?", id)
	err = row.Scan(&u.Id, &u.Displayname, &u.Email, &u.Admin)
	if err != nil {
		return u, err
	}

	return u, nil
}

func getUsernameById(id int) string {
	var u user
	db, err := getDbConnection()
	if err != nil {
		return "not found"
	}
	defer db.Close()
	row := db.QueryRow("SELECT Id, Displayname, Email, Admin FROM user WHERE Id = ?", id)
	err = row.Scan(&u.Id, &u.Displayname, &u.Email, &u.Admin)
	if err != nil {
		return "not found"
	}

	return u.Displayname
}

func getNewestBuildExecutions(limit int) ([]buildExecution, error) {
	var be buildExecution
	var beList []buildExecution

	db, err := getDbConnection()
	if err != nil {
		return beList, err
	}
	defer db.Close()
	query := "SELECT id, build_definition_id, action_log, result, execution_time, " +
		"executed_at FROM build_execution ORDER BY executed_at DESC"
	if limit > 0 {
		query += " LIMIT " + strconv.Itoa(limit)
	}
	rows, err := db.Query(query)
	if err != nil {
		return beList, err
	}

	for rows.Next() {
		err = rows.Scan(&be.Id, &be.BuildDefinitionId, &be.ActionLog, &be.Result, &be.ExecutionTime, &be.ExecutedAt)
		if err != nil {
			return beList, err
		}

		beList = append(beList, be)
		be = buildExecution{}
	}

	return beList, nil
}

func getNewestBuildDefinitions(limit int) ([]buildDefinition, error) {
	var bd buildDefinition
	var bdList []buildDefinition

	db, err := getDbConnection()
	if err != nil {
		return bdList, err
	}
	defer db.Close()
	query := "SELECT id, altered_by, caption, enabled, deployment_enabled, repo_hoster, repo_hoster_url, " +
		"repo_fullname, repo_username, repo_secret, repo_branch, altered_at FROM build_definition ORDER BY altered_at DESC"
	if limit > 0 {
		query += " LIMIT " + strconv.Itoa(limit)
	}
	rows, err := db.Query(query)
	if err != nil {
		return bdList, err
	}

	for rows.Next() {
		err = rows.Scan(&bd.Id, &bd.AlteredBy, &bd.Caption, &bd.Enabled, &bd.DeploymentEnabled, &bd.RepoHoster, &bd.RepoHosterUrl,
			&bd.RepoFullname, &bd.RepoUsername, &bd.RepoSecret, &bd.RepoBranch, &bd.AlteredAt)
		if err != nil {
			return bdList, err
		}

		bdList = append(bdList, bd)
		bd = buildDefinition{}
	}

	return bdList, nil
}