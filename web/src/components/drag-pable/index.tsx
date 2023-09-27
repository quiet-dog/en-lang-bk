import { Component, ReactNode } from "react";
import RGL, { WidthProvider, Responsive, Layout } from "react-grid-layout";
import _ from "lodash";
import { Props, State } from "./types";
import "./index.css";
import { Card, CardHeader } from "@nextui-org/card";
import { Link} from "@nextui-org/link"

export type CompactType = "horizontal" | "vertical" | null;

const ResponsiveReactGridLayout = WidthProvider(Responsive);

const availableHandles = ["s", "w", "e", "n", "sw", "nw", "se", "ne"];

export class DragPable extends Component<Props, State> {
  static defaultProps: Props = {
    className: "layout",
    rowHeight: 30,
    onLayoutChange: function () {},
    cols: { lg: 30, md: 10, sm: 6, xs: 4, xxs: 2 },
  };

  state: State = {
    currentBreakpoint: "lg",
    compactType: "vertical",
    resizeHandles: ["se"],
    mounted: false,
    layouts: { lg: generateLayout(["s", "w", "e", "n", "sw", "nw", "se"]) },
  };

  componentDidMount() {
    this.setState({ mounted: true });
  }

  generateDOM(): JSX.Element[] {
    return _.map(this.state.layouts.lg, function (l, i) {
      return (
        <div key={i}>
          <Card style={{ width: "100%", height: "100%" }}>
            <h1 className="text-2xl">title</h1>
            <h3>描述信息</h3>
            <Link href="#">链接地址</Link>
          </Card>
        </div>
      );
    });
  }

  onBreakpointChange: (Breakpoint: string) => void = (breakpoint) => {
    this.setState({
      currentBreakpoint: breakpoint,
    });
  };

  onCompactTypeChange: () => void = () => {
    const { compactType: oldCompactType } = this.state;
    const compactType =
      oldCompactType === "horizontal"
        ? "vertical"
        : oldCompactType === "vertical"
        ? null
        : "horizontal";
    this.setState({ compactType });
  };

  onResizeTypeChange: () => void = () => {
    const resizeHandles =
      this.state.resizeHandles === availableHandles ? ["se"] : availableHandles;
    this.setState({
      resizeHandles,
      layouts: { lg: generateLayout(resizeHandles) },
    });
  };

  // onLayoutChange(layout: RGL.Layout[], layouts: RGL.Layouts) {
  //   this.props.onLayoutChange(layout, layouts);
  // };

  onLayoutChange: (layout: RGL.Layout[], layouts: RGL.Layouts) => void = (
    layout,
    layouts
  ) => {
    this.props.onLayoutChange(layout, layouts);
  };

  // onNewLayout: EventHandler = () => {
  //   this.setState({
  //     layouts: { lg: generateLayout(this.state.resizeHandles) }
  //   });
  // };

  onDrop: (layout: Layout[], item: Layout, e: Event) => void = (elemParams) => {
    alert(`Element parameters: ${JSON.stringify(elemParams)}`);
  };

  render(): ReactNode {
    return (
      <>
        <div
          style={{
            position: "relative",
          }}
        >
          <ResponsiveReactGridLayout
            {...this.props}
            layouts={this.state.layouts}
            onBreakpointChange={this.onBreakpointChange}
            onLayoutChange={this.onLayoutChange}
            onDrop={this.onDrop}
            // WidthProvider option
            measureBeforeMount={false}
            // I like to have it animate on mount. If you don't, delete `useCSSTransforms` (it's default `true`)
            // and set `measureBeforeMount={true}`.
            useCSSTransforms={this.state.mounted}
            compactType={this.state.compactType}
            preventCollision={!this.state.compactType}
            margin={[20, 10]}
          >
            {this.generateDOM()}
          </ResponsiveReactGridLayout>
        </div>
      </>
    );
  }
}

function generateLayout(resizeHandles: string[]): Layout[] {
  // return _.map(_.range(0, 25), function(item, i) {
  //   var y = Math.ceil(Math.random() * 4) + 1;
  //   return {
  //     x: Math.round(Math.random() * 5) * 2,
  //     y: Math.floor(i / 6) * y,
  //     w: 2,
  //     h: y,
  //     i: i.toString(),
  //     static: Math.random() < 0.05,
  //     resizeHandles
  //   };
  // }) as Layout[];

  return [
    { x: 1, y: 1, w: 8, h: 10, i: "0", static: false, resizeHandles },
    { x: 9, y: 1, w: 8, h: 10, i: "1", static: false, resizeHandles },
    { x: 17, y: 1, w: 8, h: 10, i: "2", static: false, resizeHandles },
  ] as Layout[];
}
