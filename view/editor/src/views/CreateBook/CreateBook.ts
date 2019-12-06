import Vue from "vue"
import Component from "vue-class-component"
import Book from "@/services/ControllerBook"
import session, { Session } from "@/services/ServiceSession"
import * as route from "@/router"
import navigation from "@/services/ServiceNavigation"

@Component
export default class CreateBook extends Vue {
	session:Session = session
	validForm:boolean = false
	title:string = ""
	genre:string = ""
	description:string = ""
	titleRules:any = [(v:string) => !!v || "Title is required"]
	genreRules:any = [(v:string) => !!v || "Genre is required"]
	descriptionRules:any = [(v:string) => !!v || "Description is required"]
	genres:string[] = [
		"Roman",
		"Science-fiction",
		"Fantasy",
		"BD",
		"Conte",
	]

	cancel() {
		navigation.replaceView(route.resumePagePath)
	}

	create() {
		const b:Book = new Book(this.title, this.description, this.genre)
		b.record(this.session.user.getToken(), () => {
			this.session.addBook(b)
			navigation.replaceView(route.buildBookPagePath(b.identifier))
		})
	}
}
