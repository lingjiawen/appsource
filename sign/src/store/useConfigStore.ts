import {defineStore} from "pinia";
import {getAboutApi, getConfigApi, getHelpApi} from "@/api/mock";

interface Config {
    list: {
        name: string;
        contact: string;
        contactUrl: string;
        notice: string;
        popupHtml: string;
        buyUrl: string;
        freeOpen: string;
        icp: string;
        vats: string;
        security: string;
    };
}

export const useConfigStore = defineStore("config", {
    state: () => ({
        navbarTitle: '',
        contactName: '',
        contactUrl: '',
        notice: '',
        popupHtml: '',
        buyUrl: '',
        freeOpen: false,
        icp: '',
        vats: '',
        security: '',
        loading: true,

        helpList: [],
        aboutList: [],
    }),
    actions: {
        setConfig(config: Config) {
            this.navbarTitle = config.list.name;
            this.contactName = config.list.contact;
            this.contactUrl = config.list.contactUrl;
            this.notice = config.list.notice;
            this.popupHtml = config.list.popupHtml;
            this.buyUrl = config.list.buyUrl;
            this.freeOpen = config.list.freeOpen === '1';
            this.icp = config.list.icp;
            this.vats = config.list.vats;
            this.security = config.list.security;
        },
        setLoading(loading: boolean) {
            this.loading = loading;
        },
        setHelpList(helpList: any[]) {
            this.helpList = helpList;
        },
        setAboutList(aboutList: any[]) {
            this.aboutList = aboutList;
        },
        async fetchConfig() {
            /*
             buyUrl: "https://baidu.com"
             contact: "联系客服"
             contactUrl: "https://baidu.com"
             freeOpen: "0"
             freeP12: ""
             freePassword: ""
             freeProvision: ""
             iconDarkUrl: "https://img0.baidu.com/it/u=104684862,1159259552&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=519"
             iconLightUrl: "https://img0.baidu.com/it/u=104684862,1159259552&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=519"
             icp: "粤ICP备xxxxxxxxxx号"
             name: "芒果签"
             notice: "芒果签是一款IOS免费的签名工具，更方便提供给iOS开发者的bug内测，用于签名测试安装app."
             popupHtml: "<p>弹窗文字1</p><p>弹窗文字2</p>"
             security: "粤公网安备xxxxxxxxxxxxxx号"
             vats: "粤B2-xxxxxx"
            */
            try {
                const configResult = await getConfigApi(); // 获取配置
                this.setConfig(configResult);
            } catch (error) {
                console.error('获取配置失败', error);
            }
        },
        async fetchHelp() {
            try {
                const helpResult = await getHelpApi(); // 获取配置
                this.setHelpList(helpResult.list);
            } catch (error) {
                console.error('获取配置失败', error);
            }
        },
        async fetchAbout() {
            try {
                const aboutResult = await getAboutApi(); // 获取配置
                this.setAboutList(aboutResult.list);
            } catch (error) {
                console.error('获取配置失败', error);
            }
        }
    },
    getters: {
        getNavbarTitle: (state) => state.navbarTitle,
        getContactName: (state) => state.contactName,
        getContactUrl: (state) => state.contactUrl,
        getNotice: (state) => state.notice,
        getPopupHtml: (state) => state.popupHtml,
        getBuyUrl: (state) => state.buyUrl,
        isFreeOpen: (state) => state.freeOpen,
        getIcp: (state) => state.icp,
        getVats: (state) => state.vats,
        getSecurity: (state) => state.security,

        getHelpList: (state) => state.helpList,
        getAboutList: (state) => state.aboutList,
    },
});
