import axios from "axios"

/**
 * Override axios
 */
class Http {
	private _url:string = process.env.VUE_APP_SERVER_URL;
	private _port:string = process.env.VUE_APP_SERVER_PORT;

	getURL():string {
		return `${this._url}:${this._port}/`
	}

	getDefaultHeaders():any {
		return { "Content-Type": "application/json" }
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
		callbackSuccess:(response: any) => void, callbackError:(error: any) => void):void {
		axios.get(this.getURL() + url, { headers }).then(callbackSuccess).catch(callbackError)
	}

	post(url:string, data:any, headers:any,
		callbackSuccess:(response: any) => void, callbackError:(error: any) => void):void {
		axios.post(this.getURL() + url, data, { headers }).then(callbackSuccess).catch(callbackError)
	}

	put(url:string, data:any, headers:any,
		callbackSuccess:(response: any) => void, callbackError:(error: any) => void):void {
		axios.put(this.getURL() + url, data, { headers }).then(callbackSuccess).catch(callbackError)
	}

	delete(url:string, headers:any,
		callbackSuccess:(response: any) => void, callbackError:(error: any) => void):void {
		axios.delete(this.getURL() + url, { headers }).then(callbackSuccess).catch(callbackError)
	}
}

const httpService = new Http()

export default httpService
