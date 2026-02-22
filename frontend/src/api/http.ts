import axios from "axios";

const http = axios.create({
  baseURL: "/",
  timeout: 15000,
});

http.interceptors.response.use(
  (response) => response,
  (error) => {
    const msg =
      error?.response?.data?.status_message ||
      error?.response?.data?.error ||
      error?.message ||
      "请求失败";
    return Promise.reject(new Error(msg));
  },
);

export default http;
