import { Component } from "react";
import { type DraggableProps } from "react-beautiful-dnd";

export type DragNodeProps = Omit<DraggableProps, "children"> & {
    title: string;
    // 描述
    description?: string;
    // 备注
    remark?: string;
}