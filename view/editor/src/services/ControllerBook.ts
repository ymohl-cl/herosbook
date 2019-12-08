import httpService from "@/services/ServiceHttp"

export default class Book {
	identifier: string = ""
	title: string
	description: string
	genre: string
	publish: boolean = false
	owner: string = ""
	nodeIds: string[] = []
	creationDate: Date = new Date(0)
	categoriesID: string[] = []

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
	}

	get(identifier:string, userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.get(`api/books/${identifier}`, headers, (resp:any) => {
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

		httpService.post("api/books", this, headers, (resp: any) => {
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

		httpService.put("api/books", this, headers, (resp: any) => {
			this.unmarshall(resp.data)
			callbackSuccess()
		}, (error:any) => {
			console.log("error")
			console.log(error)
		})
	}

	delete(userToken:string, callbackSuccess: () => void) {
		const headers = httpService.appendHeaders(httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${userToken}`)

		httpService.delete(`api/books/${this.identifier}`, headers, (resp:any) => {
			callbackSuccess()
		}, (error:any) => {
			console.log("error")
			console.log(error)
		})
	}
}