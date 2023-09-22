import App from "@/App";
import { Root } from "@/views/root";
import { createBrowserRouter } from "react-router-dom";

const router = createBrowserRouter([
	{
		path: "/",
		element: <App />,
		children: [
			{
				path: "/",
				element: <Root />,
			},
			{
				path: "/home",
			},
		],
	},
]);

export default router;
