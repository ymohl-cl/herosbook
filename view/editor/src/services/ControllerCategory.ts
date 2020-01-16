import { AxiosResponse } from "axios"
import httpService from "@/services/ServiceHttp"

export class Category {
	bookId:string = ""
	identifier: string = ""
	title: string = ""
	description: string = ""
	type: string = ""

/*	constructor(bookId:string, title: string, description: string, type: string) {
		this.bookId = bookId
		this.title = title
		this.description = description
		this.type = type
	}*/

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

export default class Categories {
	persons: Category[] = []
	locations: Category[] = []
	customs: Category[] = []

	deleteCategory(userToken:string, cat: Category, callbackSuccess:() => void) {
		let index = 0
		if (cat.type == "person") {
			index = this.persons.findIndex(category => category.identifier === cat.identifier)
			if (index >= 0) {
				this.persons[index].delete(userToken, () => {
					this.persons.splice(index, 1)
					callbackSuccess()
				})
			}
		} else if (cat.type == "location") {
			index = this.locations.findIndex(category => category.identifier === cat.identifier)
			if (index >= 0) {
				this.locations[index].delete(userToken, () => {
					this.locations.splice(index, 1)
					callbackSuccess()
				})
			}
		} else {
			index = this.customs.findIndex(category => category.identifier === cat.identifier)
			if (index >= 0) {
				this.customs[index].delete(userToken, () => {
					this.customs.splice(index, 1)
					callbackSuccess()
				})
			}
		}
	}

	addCategory(cat: Category) {
		if (cat.type == "person") {
			this.persons.push(cat)
		} else if (cat.type == "location") {
			this.locations.push(cat)
		} else {
			this.customs.push(cat)
		}
	}
}
