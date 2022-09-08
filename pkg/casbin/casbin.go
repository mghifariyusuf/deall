package casbin

import (
	"database/sql"
	"deall/cmd/lib/customError"
	"encoding/json"
	"log"
	"os"

	sqladapter "github.com/Blank-Xu/sql-adapter"
	"github.com/sirupsen/logrus"

	"github.com/casbin/casbin/v2"
)

var enforcer *casbin.Enforcer

type RoleData struct {
	Role   string `json:"role"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type CasbinConfig struct {
	db *sql.DB
}

func New(db *sql.DB) *CasbinConfig {
	return &CasbinConfig{db: db}
}

func (c *CasbinConfig) Register() error {
	var role []*RoleData
	adapter, err := sqladapter.NewAdapter(c.db, "mysql", "casbin_rule")
	if err != nil {
		logrus.Error(err)
		return err
	}

	enforcer, err = casbin.NewEnforcer("cmd/config/rbac_config.conf", adapter)
	if err != nil {
		logrus.Error(err)
		return err
	}

	if err = enforcer.LoadPolicy(); err != nil {
		logrus.Error(err)
		return err
	}

	data, err := os.ReadFile("cmd/config/casbin.json")
	if err != nil {
		logrus.Error(err)
		return err
	}

	err = json.Unmarshal(data, &role)
	if err != nil {
		logrus.Error(err)
		return err
	}

	for _, value := range role {
		err = InsertPolicy(value)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func CheckPolicy(data *RoleData) (bool, error) {
	isAuthorize, err := enforcer.Enforce(data.Role, data.Path, data.Method)
	if err != nil {
		logrus.Error(err)
		return false, customError.ErrNotAuthorize
	}

	if isAuthorize {
		return true, nil
	}
	return false, customError.ErrNotAuthorize
}

func InsertPolicy(data *RoleData) error {
	_, err := enforcer.AddPolicy(data.Role, data.Path, data.Method)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
