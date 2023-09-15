import { Card } from "antd";
import { NavCardProps } from "./types";
import Meta from "antd/es/card/Meta";

export default function NavCard(props: NavCardProps) {
	return (
		<>
			<Card cover={<img src={props.imgUrl} onClick={props.onClick} />}>
				<Meta title={props.title} />
			</Card>
		</>
	);
}
