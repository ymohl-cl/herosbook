import httpService from "@/services/ServiceHttp"
import Category from "@/services/ControllerCategory"

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
