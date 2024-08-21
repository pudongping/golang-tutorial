package snow_flake

import (
	"fmt"
	"reflect"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/sony/sonyflake"
)

func SnowFlake1() {
	var (
		node *snowflake.Node
		st   time.Time
		err  error
	)

	startTime := "2024-08-20" // 初始化一个开始的时间，表示从这个时间开始算起
	machineID := 1            // 机器 ID

	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}

	snowflake.Epoch = st.UnixNano() / 1000000
	// 根据指定的开始时间和机器ID，生成节点实例
	node, err = snowflake.NewNode(int64(machineID))
	if err != nil {
		panic(err)
	}

	// 生成并输出 ID
	id := node.Generate()

	fmt.Printf("Int64  ID: %d type of: %T -> Type %v -> Value %v \n", id, id, reflect.TypeOf(id), reflect.ValueOf(id))
	fmt.Printf("Int64  ID: %d\n", id.Int64()) // 也可以直接调用 Int64() 方法
	fmt.Printf("String ID: %s\n", id)
	fmt.Printf("Base2  ID: %s\n", id.Base2())
	fmt.Printf("Base64 ID: %s\n", id.Base64())
}

func SnowFlake2() {
	var (
		sonyFlake     *sonyflake.Sonyflake
		sonyMachineID uint16
		st            time.Time
		err           error
	)

	startTime := "2024-08-20" // 初始化一个开始的时间，表示从这个时间开始算起
	machineID := 1            // 机器 ID
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		panic(err)
	}

	sonyMachineID = uint16(machineID)
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: func() (uint16, error) { return sonyMachineID, nil },
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	if sonyFlake == nil {
		panic("sonyflake not created")
	}

	id, err := sonyFlake.NextID()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Int64  ID: %d type of: %T -> Type %v -> Value %v \n", id, id, reflect.TypeOf(id), reflect.ValueOf(id))

}
