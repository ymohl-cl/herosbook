import Vue from "vue"
import Component from "vue-class-component"

export const WarningMessage:string = "warning"
export const InfoMessage:string = "info"
export const ErrorMessage:string = "error"
export const SucessMessage:string = "success"
const timeoutDuration = 10000

export class Message {
	message: string
	type: string

	constructor(message: string, type: string) {
		this.message = message
		this.type = type
	}
}

@Component
export default class AlertMessage extends Vue {
	current: Message = new Message("init-alerting", InfoMessage)
	showAlert: boolean = false
	listMessage: Message[] = []

	addAlert(message: string, type: string) {
		console.log("new alert")
		console.log(message)
		console.log(type)
		console.log(this.current.message)
		this.current = new Message(message, type)
		this.showAlert = true
		setTimeout(() => {
			this.showAlert = false
			this.listMessage.push(this.current)
			console.log("close message")
		}, timeoutDuration)
	}

	removeMessage(id: number) {
		this.listMessage.splice(id, 1)
	}
}
