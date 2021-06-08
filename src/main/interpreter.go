package main

import (
	"fmt"
	"instructions"
	"instructions/base"
	"rtda"
	"rtda/heap"
)

func interpret(methodInfo *heap.Method) {
	//codeAttr := methodInfo.CodeAttribute()
	//maxLocals := codeAttr.MaxLocals()
	//maxStack := codeAttr.MaxStack()
	//bytecode := codeAttr.Code()
	//
	//thread := rtda.NewThread()
	//frame := thread.NewFrame(maxLocals, maxStack)
	//thread.PushFrame(frame)
	//
	//defer catchErr(frame)
	//loop(thread, bytecode)
	thread := rtda.NewThread()
	frame := thread.NewFrame(methodInfo)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, methodInfo.Code())
}

func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}

	for {
		//执行当前栈帧，获取指令计数器
		pc := frame.NextPC()
		thread.SetPC(pc)

		// 解码译码decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// 指令执行execute
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
