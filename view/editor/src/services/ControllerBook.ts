import { AxiosResponse } from "axios"
import httpService from "@/services/ServiceHttp"
import Categories, * as Category from '@/services/ControllerCategory'

export function getBooks(token: string, callbackSuccess: (jsonData: any) => void) {
	const headers = httpService.appendHeaders(
		httpService.getDefaultHeaders(),
		"Authorization", `Bearer ${token}`,
	)
	httpService.post("api/books/_searches", {}, headers, (resp: AxiosResponse) => {
		callbackSuccess(resp.data)
	})
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
	categories: Categories = new Categories()

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
		if (json.categories.persons !== null) {
			for (let i = 0; i < json.categories.persons.length; i += 1) {
				let category = new Category.Category()
				category.unmarshall(json.categories[i])
				this.categories.persons.push(category)
			}
		}
		if (json.categories.locations !== null) {
			for (let i = 0; i < json.categories.locations.length; i += 1) {
				let category = new Category.Category()
				category.unmarshall(json.categories[i])
				this.categories.locations.push(category)
			}
		}
		if (json.categories.customs !== null) {
			for (let i = 0; i < json.categories.customs.length; i += 1) {
				let category = new Category.Category()
				category.unmarshall(json.categories[i])
				this.categories.customs.push(category)
			}
		}
		console.log(json.categories.persons.length)
		console.log(json.categories.locations.length)
		console.log(json.categories.customs.length)
	}

	get(identifier:string, userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		console.log("before get request")
		httpService.get(`api/books/${identifier}`, headers, (resp: AxiosResponse) => {
			this.unmarshall(resp.data)
			console.log("hello")
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
}
