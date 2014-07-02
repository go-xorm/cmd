package main

import (
	"errors"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	_ "github.com/go-xorm/ql"
	"github.com/go-xorm/xorm"
	"github.com/go-xweb/xweb"
	_ "github.com/lunny/ql/driver"
)

var CmdWeb = &Command{
	UsageLine: "web",
	Short:     "web",
	Long: `
`,
}

func init() {
	CmdWeb.Run = runWeb
	CmdWeb.Flags = map[string]bool{}
}

var (
	orm *xorm.Engine
)

type Engine struct {
	Id         int64
	Name       string `xorm:"unique"`
	Driver     string
	DataSource string
}

type MainAction struct {
	*xweb.Action

	home xweb.Mapper `xweb:"/"`
	addb xweb.Mapper
	del  xweb.Mapper
	view xweb.Mapper
}

func (c *MainAction) Home() error {
	engines := make([]Engine, 0)
	err := orm.Find(&engines)
	if err != nil {
		return err
	}
	return c.Render("root.html", &xweb.T{
		"engines": engines,
		"tables":  []core.Table{},
		"records": [][]string{},
		"columns": []string{},
		"id":      0,
	})
}

func (c *MainAction) Addb() error {
	if c.Method() == "GET" {
		return c.Render("add.html")
	}

	var engine Engine
	err := c.MapForm(&engine)
	if err != nil {
		return err
	}
	_, err = orm.Insert(&engine)
	if err != nil {
		return err
	}
	return c.Render("addsuccess.html")
}

func (c *MainAction) Del() error {
	id, err := c.GetInt("id")
	if err != nil {
		return err
	}

	_, err = orm.Id(id).Delete(new(Engine))
	if err != nil {
		return err
	}

	return c.Render("delsuccess.html")
}

var (
	ormCache  = make(map[string]*xorm.Engine)
	cacheLock sync.RWMutex
)

func getOrm(name string) *xorm.Engine {
	cacheLock.RLock()
	defer cacheLock.RUnlock()
	if o, ok := ormCache[name]; ok {
		return o
	}
	return nil
}

func setOrm(name string, o *xorm.Engine) {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	ormCache[name] = o
}

func (c *MainAction) View() error {
	id, err := c.GetInt("id")
	if err != nil {
		return err
	}

	engine := new(Engine)
	has, err := orm.Id(id).Get(engine)
	if err != nil {
		return err
	}

	if !has {
		return errors.New("db is not exist")
	}

	o := getOrm(engine.Name)

	if o == nil {
		o, err = xorm.NewEngine(engine.Driver, engine.DataSource)
		if err != nil {
			return err
		}

		setOrm(engine.Name, o)
	}

	tables, err := o.DBMetas()
	if err != nil {
		return err
	}

	engines := make([]Engine, 0)
	err = orm.Find(&engines)
	if err != nil {
		return err
	}

	var records = make([][]*string, 0)
	var columns = make([]string, 0)
	tb := c.GetString("tb")
	if tb != "" {
		rows, err := o.DB().Query("select * from " + tb)
		if err != nil {
			return err
		}
		defer rows.Close()

		columns, err = rows.Columns()
		if err != nil {
			return err
		}

		for rows.Next() {
			datas := make([]*string, len(columns))
			err = rows.ScanSlice(&datas)
			if err != nil {
				return err
			}
			records = append(records, datas)
		}
	}

	return c.Render("root.html", &xweb.T{
		"engines": engines,
		"tables":  tables,
		"records": records,
		"columns": columns,
		"id":      id,
	})
}

func runWeb(cmd *Command, args []string) {
	var err error
	orm, err = xorm.NewEngine("ql", "./xorm.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = orm.Sync(&Engine{})
	if err != nil {
		fmt.Println(err)
		return
	}

	xweb.AddAction(&MainAction{})
	xweb.Run(":8989")
}
