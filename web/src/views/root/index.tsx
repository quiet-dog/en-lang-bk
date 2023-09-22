import { DragMenu } from "@/components/drag-menu";
import { Component, ReactNode } from "react";
import { DragPable } from "@/components/drag-pable";

export class Root extends Component {
  constructor(props: any) {
    super(props);
  }

  render(): ReactNode {
    return (
      <>
        <Common>
          <DragMenu />
          <DragPable />
        </Common>
      </>
    );
  }
}
