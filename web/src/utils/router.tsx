import { useNavigate, type NavigateOptions } from "react-router-dom";

export function goUrl(to: string, options?: NavigateOptions | undefined) {
	const navigate = useNavigate();
	navigate(to, options);
}
