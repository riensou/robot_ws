package db_writer

import (
	_ "github.com/aler9/goroslib"
	_ "github.com/gwaxG/robot_ws/backend/pkg/common"
	database2 "github.com/gwaxG/robot_ws/backend/pkg/database"
	"github.com/gwaxG/robot_ws/monitor/pkg/structs"
	"log"
)

type Core struct {
	ros         Ros
	database    database2.DataBase
	analyticsCh chan structs.RolloutAnalytics
}

func (c *Core) Init() {
	c.analyticsCh = make(chan structs.RolloutAnalytics)
	c.ros = Ros{}
	c.ros.Init(c.analyticsCh)
	c.database = database2.DataBase{}
	c.database.Init()
}

func (c *Core) Start() {
	// Closing connections on error
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in Core.Start()", r)
			c.Close()
		}
	}()

	var analytics structs.RolloutAnalytics

	for {
		select {
			// New analytics coming from ROS
			case analytics = <- c.analyticsCh:
				go c.database.AddNewRolloutAnalytics(analytics)
		}
	}
}

func (c *Core) Close() {
	c.ros.Close()
	c.database.Close()
}

