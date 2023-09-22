import { Component, ReactNode } from "react";
import { Draggable } from "react-beautiful-dnd";
import { DragNodeProps } from "./types";

export class DragNode extends Component<DragNodeProps> {
  constructor(props: DragNodeProps) {
    super(props);
  }

  render(): ReactNode {
    return (
      <>
        <Draggable
          draggableId={this.props.draggableId}
          index={this.props.index}
        >
          {(provided, _snapshot) => (
            <div
              ref={provided.innerRef}
              {...provided.draggableProps}
              {...provided.dragHandleProps}
            >
              <h1>{this.props.title}</h1>
              <h2>{this.props.description}</h2>
              <h3>{this.props.remark}</h3>
            </div>
          )}
        </Draggable>
      </>
    );
  }
}
