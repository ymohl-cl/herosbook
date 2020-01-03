import Book, * as Books from "@/services/ControllerBook"
import User from "@/services/ControllerUser"
import AlertMessage from "@/components/AlertMessage/AlertMessage"
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
		this.books.push(b)
	}

	getBooks() {
		this.books = Books.getBooks(this.user.getToken())
	}
}

const session = new Session()

export default session
