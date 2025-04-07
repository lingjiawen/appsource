import { defineStore } from 'pinia';
import {AppDownloadNoticeMessage, AppDownloadNoticeStates} from "/@/stores/interface";

export const appDownloadStore = defineStore({
    id: 'appDownloadStore',
    state: (): AppDownloadNoticeStates => <AppDownloadNoticeStates>({
        message: {},
        socketStatus: false
    }),
    getters: {
        getMessages(): AppDownloadNoticeMessage {
            return this.message;
        },
        getSocketStatus(): boolean {
            return this.socketStatus;
        }
    },
    actions: {
        setMessages(messages: AppDownloadNoticeMessage) {
            this.message = messages;
        },
        setSocketStatus(status: boolean) {
            this.socketStatus = status;
        }
    },
});
