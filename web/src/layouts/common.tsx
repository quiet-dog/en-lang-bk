import React, { ReactNode } from "react";

import {
  Flexboard,
  FlexboardProvider,
  FlexboardFrame,
  ResizerType,
  Position,
} from "@dorbus/flexboard";
export class Common extends React.Component<{
  children?: ReactNode[];
  title?: string;
}> {
  render() {
    return (
      <>
        <FlexboardProvider>
        <FlexboardFrame>
            <div>{this.props.children![0]}</div>
          </FlexboardFrame>
          <Flexboard
            direction={Position.right}
            draggable={true}
            width={220}
            minWidth={200}
            maxWidth={1300}
            resizerType={ResizerType.line}
          >
            <div style={{ width: "100%", position: "relative",zIndex:3 }}>
                {this.props.children![1]}
            </div>
          </Flexboard>
       
        
        </FlexboardProvider>
        {/* <Layout>
          <Sider style={{ background: "white" }}>
            {this.props.children![0]}
          </Sider>
          <Content>{this.props.children![1]}</Content>
        </Layout> */}
      </>
    );
  }
}
