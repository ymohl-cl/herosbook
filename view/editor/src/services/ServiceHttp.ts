import axios, { AxiosResponse } from "axios"
import AlertMessage, * as Alert from '@/components/AlertMessage/AlertMessage.ts';

class ErrorResponse {
	message:string = ""
	id:string = ""

	constructor(json:any) {
		this.message = json.message
		this.id = json.identifier
	}
}

/**
 * Override axios
 */
class Http {
	private _url:string = process.env.VUE_APP_SERVER_URL
	private _port:string = process.env.VUE_APP_SERVER_PORT
	private _alert:AlertMessage = new AlertMessage()

	getURL():string {
		return `${this._url}:${this._port}/`
	}

	getDefaultHeaders():any {
		return { "Content-Type": "application/json" }
	}

	setAlertMessage(alert:AlertMessage):void {
		this._alert = alert
	}

	sendAlert(resp:AxiosResponse):void {
		console.log(resp)
		const err = new ErrorResponse(resp.data)
		const message = "["+err.id+"]: " + err.message
		if (resp.status >= 400 && resp.status < 500) {
			this._alert.addAlert(message, Alert.WarningMessage)
		} else {
			this._alert.addAlert(message, Alert.ErrorMessage)
		}
	}

	appendHeaders(headers:any, key:string, value:string):any {
		const newHeaders = headers

		newHeaders[key] = value
		return newHeaders
	}

	/*
	** method http
	*/
	get(url:string, headers:any,
		callbackSuccess:(resp: AxiosResponse) => void):void {
		axios.get(this.getURL() + url, { headers }).then(
			response => callbackSuccess(response)).catch(
			error => this.sendAlert(error.response))
	}

	post(url:string, data:any, headers:any,
		callbackSuccess:(resp: AxiosResponse) => void):void {
		axios.post(this.getURL() + url, data, { headers }).then(
			response => callbackSuccess(response)).catch(
			error => this.sendAlert(error.response))
	}

	put(url:string, data:any, headers:any,
		callbackSuccess:(resp: AxiosResponse) => void):void {
		axios.put(this.getURL() + url, data, { headers }).then(
			response => callbackSuccess(response)).catch(
			error => this.sendAlert(error.response))
	}

	delete(url:string, headers:any,
		callbackSuccess:(resp: AxiosResponse) => void):void {
		axios.delete(this.getURL() + url, { headers }).then(
			response => callbackSuccess(response)).catch(
			error => this.sendAlert(error.response))
	}
}

const httpService = new Http()

export default httpService
