import httpService from "@/services/ServiceHttp"

export default class Node {
	bookId:string = "";
	identifier: string = "";
	title: string;
	description: string;
	content: string = "";
	categories: any[] = []; // TODO category

	constructor(bookId:string, title: string, description: string) {
		this.bookId = bookId
		this.title = title
		this.description = description
	}

	unmarshall(json: any) {
		this.identifier = json.identifier
		this.title = json.title
		this.description = json.description
		this.content = json.content
		// TODO ERROR json.categories null
		if (json.categories !== null) {
			this.categories = json.categories
		} else {
			this.categories = []
		}
	}

	addCategory(catIdentifier:string) {
		const index = this.categories.findIndex(id => id === catIdentifier)

		if (index === -1) {
			this.categories.push(catIdentifier)
		}
	}

	removeCategory(catIdentifier:string) {
		const index = this.categories.findIndex(id => id === catIdentifier)

		this.removeCategoryByIndex(index)
	}

	removeCategoryByIndex(index:number) {
		return this.categories.splice(index, 1)
	}

	get(identifier:string, userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.get(`api/books/${this.bookId}/node/${identifier}`, headers, (resp:any) => {
			this.unmarshall(resp.data)
			callbackSuccess()
		}, (error:any) => {
			console.log("error")
			console.log(error)
		})
	}

	record(userToken: string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.post(`api/books/${this.bookId}/node`, this, headers, (resp: any) => {
			this.unmarshall(resp.data)
			callbackSuccess()
		}, (error:any) => {
			console.log("error")
			console.log(error)
		})
	}

	update(userToken: string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.put(`api/books/${this.bookId}/node`, this, headers, (resp: any) => {
			// this.unmarshall(resp.data)
			callbackSuccess()
		}, (error:any) => {
			console.log("error")
			console.log(error)
		})
	}

	delete(userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.delete(`api/books/${this.bookId}/node/${this.identifier}`, headers, (resp:any) => {
			callbackSuccess()
		}, (error:any) => {
			console.log("error")
			console.log(error)
		})
	}
}
