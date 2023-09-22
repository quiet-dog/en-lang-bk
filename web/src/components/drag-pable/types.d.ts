import RGL from "react-grid-layout";
export type DragPableProps = {
    className?: string,
    items: number,
    rowHeight?: number,
    onLayoutChange: (layout: RGL.Layout[]) => void,
    cols?: number,
}

export type DragPableState = {
    layout?: RGL.Layout[],
}

export type Props = {
    className: string;
    cols: { [key: string]: number };
    onLayoutChange: Function;
    rowHeight: number;
};
export type State = {
    currentBreakpoint: string;
    compactType: CompactType;
    mounted: boolean;
    resizeHandles: string[];
    layouts: { [key: string]: Layout[] };
};