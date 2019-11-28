import Book from '@/services/ControllerBook'
import User from '@/services/ControllerUser'
import AlertMessage from '@/components/AlertMessage/AlertMessage'

export class Session {
	books: Book[] = []
	user: User = new User()
	alert: AlertMessage = new AlertMessage()

	removeBook(identifier: string) {
		var index = this.books.findIndex(element => element.identifier == identifier)
		this.books[index].delete(this.user.getToken(), () => {
			this.books.splice(index, 1)
		})
	}
	addBook(b: Book) {
		b.record(this.user.getToken(), () => {
			this.books.push(b)
		})
	}
}

let session = new Session()
export default session
