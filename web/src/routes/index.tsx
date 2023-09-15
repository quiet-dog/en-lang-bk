import App from "@/App";
import { createBrowserRouter } from "react-router-dom";
import Navigation from "@/layouts/navigation";
const router = createBrowserRouter([
	{
		path: "/",
		element: <App />,
		children: [
			{
				path: "/",
				element: <Navigation />,
			},
		],
	},
]);

export default router;
