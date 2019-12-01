import Vue from "vue"
import Component from "vue-class-component"
import navigation from "@/services/ServiceNavigation"
import session, { Session } from "@/services/ServiceSession"
import * as route from "@/router"
import ConfirmDialog from "@/components/ConfirmDialog/ConfirmDialog.vue"

@Component({
	components: { ConfirmDialog },
})
export default class ResumeHome extends Vue {
	session: Session = session
	deleteBookId:string = ""

	mounted() {
		session.getBooks()
	}

	createBook() {
		navigation.replaceView(route.createBookPagePath)
	}

	showBook(identifier: string) {
		navigation.changeView(route.buildBookPagePath(identifier))
	}

	deleteBook(idenitifer: string) {
		session.removeBook(idenitifer)
	}
}
