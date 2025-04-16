import { http } from "@/utils/http";

type ListResult = {
  list: any;
};

export function getConfigApi(): Promise<ListResult> {
  return http.request({
    url: "getconfig",
    method: "get"
  });
}

export function getHelpApi(): Promise<ListResult> {
  return http.request({
    url: "gethelp",
    method: "get"
  });
}

export function getAboutApi(): Promise<ListResult> {
  return http.request({
    url: "getabout",
    method: "get"
  });
}

export function installApi(data?: object): Promise<any> {
  return http.request({
    url: "install",
    method: "post",
    data
  });
}

export function installPrivateApi(data?: object): Promise<any> {
  return http.request({
    url: "installPrivate",
    method: "post",
    data
  });
}

export function deviceCheckApi(data?: object): Promise<any> {
  return http.request({
    url: "deviceCheck",
    method: "post",
    data
  });
}

export function getListApi(params?: object): Promise<ListResult> {
  return http.request({
    url: "/list/get",
    method: "get",
    params
  });
}

export function getListApiError(data?: object): Promise<ListResult> {
  return http.request({
    url: "/list/error",
    method: "post",
    data
  });
}
