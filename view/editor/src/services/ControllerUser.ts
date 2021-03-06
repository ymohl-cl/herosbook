import { AxiosResponse } from "axios"
import httpService from "@/services/ServiceHttp"

export default class User {
	identifier: string = ""
	pseudo: string = ""
	name: string = ""
	lastName: string = ""
	age: number = 0
	genre: string = ""
	email: string = ""

	private _connect: boolean = false
	private _token: string = ""

	login(username: string, password: string, callbackSucess: () => void) {
		const headers = httpService.appendHeaders(
			httpService.appendHeaders(httpService.getDefaultHeaders(),
				"username", username),
			"password", btoa(`${username}:${password}`),
		)

		httpService.post("login", {}, headers, (resp:AxiosResponse) => {
			this.unmarshall(resp.data)
			this._connect = true
			callbackSucess()
		})
	}

	unmarshall(json: any) {
		this.identifier = json.user.identifier
		this.pseudo = json.user.pseudo
		this.name = json.user.name
		this.lastName = json.user.lastName
		this.age = json.user.age
		this.genre = json.user.genre
		this.email = json.user.email
		this._token = json.token
	}

	record(password: string, callbackSucess: () => void) {
		console.log("me: ", this.pseudo)
		console.log("password: ", password)
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"password", btoa(`${this.pseudo}:${password}`))

		httpService.post("register", this, headers, (resp: AxiosResponse) => {
			callbackSucess()
		})
	}

	isConnected(): boolean {
		return this._connect
	}

	getToken(): string {
		return this._token
	}

	disconnect() {
		this._connect = false
		this._token = ""
	}
}
