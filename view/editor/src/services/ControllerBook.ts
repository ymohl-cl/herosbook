import { AxiosResponse } from "axios"
import httpService from "@/services/ServiceHttp"
import Category from "@/services/ControllerCategory"

export function getBooks(token: string):Book[] {
	let books: Book[] = []
	const headers = httpService.appendHeaders(
		httpService.getDefaultHeaders(),
		"Authorization", `Bearer ${token}`,
	)
	httpService.post("api/books/_searches", {}, headers, (resp: AxiosResponse) => {
		books = resp.data
	})
	return books
}

export default class Book {
	identifier: string = ""
	title: string
	description: string
	genre: string
	publish: boolean = false
	owner: string = ""
	nodeIds: string[] = []
	creationDate: Date = new Date(0)
	categories: Category[] = []

	constructor(title: string, description: string, genre: string) {
		this.title = title
		this.description = description
		this.genre = genre
	}

	unmarshall(json: any) {
		this.identifier = json.identifier
		this.publish = json.publish
		this.owner = json.owner
		this.nodeIds = json.nodeIds
		this.creationDate = json.creationDate
		this.title = json.title
		this.description = json.description
		this.genre = json.genre
		if (json.categories !== null) { // TOUPDATE
			for (let i = 0; i < json.categories.length; i += 1) {
				const category = new Category(this.identifier, "", "", "")

				category.unmarshall(json.categories[i])
				this.categories.push(category)
			}
		}
	}

	get(identifier:string, userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.get(`api/books/${identifier}`, headers, (resp:AxiosResponse) => {
			this.unmarshall(resp.data)
			callbackSuccess()
		})
	}

	record(userToken: string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.post("api/books", this, headers, (resp: AxiosResponse) => {
			this.unmarshall(resp.data)
			callbackSuccess()
		})
	}

	update(userToken: string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.put("api/books", this, headers, (resp: AxiosResponse) => {
			callbackSuccess()
		})
	}

	delete(userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.delete(`api/books/${this.identifier}`, headers, (resp:AxiosResponse) => {
			callbackSuccess()
		})
	}

	deleteCategory(userToken:string, catIdentifier:string, callbackSuccess:() => void) {
		const index = this.categories.findIndex(category => category.identifier === catIdentifier)

		if (index > -1) {
			this.categories[index].delete(userToken, () => {
				this.categories.splice(index, 1)
				callbackSuccess()
			})
		}
	}
}
