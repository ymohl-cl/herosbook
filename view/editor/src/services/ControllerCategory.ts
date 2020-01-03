import { AxiosResponse } from "axios"
import httpService from "@/services/ServiceHttp"

export default class Category {
	bookId:string = ""
	identifier: string = ""
	title: string
	description: string
	type: string = ""

	constructor(bookId:string, title: string, description: string, type: string) {
		this.bookId = bookId
		this.title = title
		this.description = description
		this.type = type
	}

	unmarshall(json: any) {
		this.identifier = json.identifier
		this.title = json.title
		this.description = json.description
		this.type = json.type
	}

	get(identifier:string, userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.get(`api/books/${this.bookId}/category/${identifier}`, headers, (resp:AxiosResponse) => {
			this.unmarshall(resp.data)
			callbackSuccess()
		})
	}

	record(userToken: string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.post(`api/books/${this.bookId}/category`, this, headers, (resp: AxiosResponse) => {
			this.unmarshall(resp.data)
			callbackSuccess()
		})
	}

	update(userToken: string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.put(`api/books/${this.bookId}/category`, this, headers, (resp: AxiosResponse) => {
			callbackSuccess()
		})
	}

	delete(userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.delete(`api/books/${this.bookId}/category/${this.identifier}`, headers, (resp:AxiosResponse) => {
			callbackSuccess()
		})
	}
}
