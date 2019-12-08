import Vue from "vue"
import Component from "vue-class-component"
import navigation from "@/services/ServiceNavigation"
import session, { Session } from "@/services/ServiceSession"
import Book from "@/services/ControllerBook"
import Node from "@/services/ControllerNode"
import TextEditor from "@/components/TextEditor/TextEditor.vue"

@Component({
	components: { TextEditor },
})
export default class ViewBook extends Vue {
	session:Session = session
	book:Book = new Book("", "", "")
	nodes:Node[] = []
	currentNode:Node = new Node("", "", "")
	editNodeMode:boolean = false;

	mounted() {
		if (!session.user.isConnected() || this.$route.params.id == null) {
			navigation.replaceView("/login")
		} else {
			this.book.get(this.$route.params.id, session.user.getToken(), () => {
				for (let i = 0; i < this.book.nodeIds.length; i += 1) {
					const node = new Node(this.book.identifier, "", "")

					this.nodes.push(node)
					node.get(this.book.nodeIds[i], session.user.getToken(), () => {
					})
				}
			})
		}
	}

	setCurrentNode(node:Node){
		this.currentNode = node
	}

	setEditNodeMode(mode:boolean){
		this.editNodeMode = mode
	}

	createNode() {
		const node = new Node(this.book.identifier, "New node", "")

		node.record(session.user.getToken(), () => {
			this.nodes.push(node)
		})
	}

	saveCurrentNode() {
		this.currentNode.update(session.user.getToken(), () => {
			console.log("save node ok")
		})
	}

	updateTextNode(content:string) {
		this.currentNode.content = content
	}
}
