import { defineStore } from 'pinia';
import {SourceImportNoticeMessage, SourceImportNoticeStates} from "/@/stores/interface";

export const sourceImportStore = defineStore({
    id: 'sourceImportStore',
    state: (): SourceImportNoticeStates => <SourceImportNoticeStates>({
        message: {},
        socketStatus: false
    }),
    getters: {
        getMessages(): SourceImportNoticeMessage {
            return this.message;
        },
        getSocketStatus(): boolean {
            return this.socketStatus;
        }
    },
    actions: {
        setMessages(messages: SourceImportNoticeMessage) {
            this.message = messages;
        },
        setSocketStatus(status: boolean) {
            this.socketStatus = status;
        }
    },
});
