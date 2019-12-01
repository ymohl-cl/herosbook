import Book from "@/services/ControllerBook"
import User from "@/services/ControllerUser"
import AlertMessage from "@/components/AlertMessage/AlertMessage.vue"
import httpService from "@/services/ServiceHttp"

export class Session {
	books: Book[] = []
	user: User = new User()
	alert: AlertMessage = new AlertMessage()

	removeBook(identifier: string) {
		const index = this.books.findIndex(element => element.identifier === identifier)

		this.books[index].delete(this.user.getToken(), () => {
			this.books.splice(index, 1)
		})
	}

	addBook(b: Book) {
		b.record(this.user.getToken(), () => {
			this.books.push(b)
		})
	}

	getBooks() {
		const headers = httpService.appendHeaders(
			httpService.getDefaultHeaders(),
			"Authorization", `Bearer ${this.user.getToken()}`,
		)

		httpService.post("api/books/_searches", {}, headers, (resp: any) => {
			this.books = resp.data
		}, (error:any) => {
			// alert error
		})
	}
}

const session = new Session()

export default session
