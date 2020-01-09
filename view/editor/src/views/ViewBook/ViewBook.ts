import Vue from "vue"
import Component from "vue-class-component"
import navigation from "@/services/ServiceNavigation"
import session, { Session } from "@/services/ServiceSession"
import Book from "@/services/ControllerBook"
import Node from "@/services/ControllerNode"
import Category from "@/services/ControllerCategory"
import TextEditor from "@/components/TextEditor/TextEditor.vue"
import ConfirmDialog from "@/components/ConfirmDialog/ConfirmDialog.vue"

@Component({
	components: {
		TextEditor,
		ConfirmDialog,
	},
})
export default class ViewBook extends Vue {
	session:Session = session
	book:Book = new Book("", "", "")
	editBookMode:boolean = false
	nodes:Node[] = []
	nodesDisplayed:Node[] = []
	currentNode:Node = new Node("", "", "")
	currentCategoryIdentifier = ""
	editNodeMode:boolean = false
	deleteNodeIdentifier:string = ""
	deleteCategoryIdentifier:string = ""
	titleNewCategory:string = ""
	categoryToAdd:string = ""

	orphanNodes:Node[] = [new Node(this.book.identifier, "title", "description")]
	treeNodes:Node[][] = [
		[new Node(this.book.identifier, "title-1", "description")],
		[
			new Node(this.book.identifier, "title-2.1", "description"),
			new Node(this.book.identifier, "title-2.2", "description"),
			new Node(this.book.identifier, "title-2.3", "description")
		],
		[new Node(this.book.identifier, "title-3", "description")],
		[new Node(this.book.identifier, "title-4", "description")],
		[new Node(this.book.identifier, "title-5", "description")],
		[new Node(this.book.identifier, "title-6", "description")],
	]
	displayConfirm:boolean = false
	callbackSuccessConfirm:any = null
	callbackCancelConfirm:any = null
	textConfirm:string = ""

	mounted() {
		if (!session.user.isConnected() || this.$route.params.id == null) {
			navigation.replaceView("/login")
		} else {
			this.book.get(this.$route.params.id, session.user.getToken(), () => {
				for (let i = 0; i < this.book.nodeIds.length; i += 1) {
					const node = new Node(this.book.identifier, "", "")

					this.nodes.push(node)
					node.get(this.book.nodeIds[i], session.user.getToken(), () => {
						this.generateDisplayNode()
					})
				}
			})
		}
	}

	generateDisplayNode() {
		this.nodesDisplayed = this.nodes.filter((node:Node) => this.currentCategoryIdentifier === ""
			|| node.categories.indexOf(this.currentCategoryIdentifier) > -1)
	}

	getCategoryByIdentifier(identifier:string) {
		return this.book.categories.find(category => category.identifier === identifier)
	}

	getCategoriesdentifier() {
		const identifiers = [{
			text: "Tous les noeuds",
			value: "",
		}]

		this.book.categories.forEach(category => identifiers.push({
			text: category.title,
			value: category.identifier,
		}))
		return identifiers
	}

	setEditBookMode(mode:boolean) {
		this.editBookMode = mode
	}

	setCurrentNode(node:Node) {
		this.currentNode = node
	}

	setEditNodeMode(mode:boolean) {
		this.editNodeMode = mode
	}

	editNode(node:Node) {
		this.setCurrentNode(node)
		this.setEditNodeMode(true)
	}

	updateTextNode(content:string) {
		this.currentNode.content = content
	}

	// TODO : put in session ?
	saveCurrentBook() {
		this.book.update(session.user.getToken(), () => {
			this.setEditBookMode(false)
		})
	}

	openAlert(displayConfirm:boolean, textConfirm:string,
		callbackSuccessConfirm: () => void, callbackCancelConfirm: () => void) {
		this.displayConfirm = displayConfirm
		this.textConfirm = textConfirm
		this.callbackSuccessConfirm = callbackSuccessConfirm
		this.callbackCancelConfirm = callbackCancelConfirm
	}

	createCategory() {
		const category = new Category(this.book.identifier, this.titleNewCategory, "Description", "person")

		category.record(session.user.getToken(), () => {
			this.book.categories.push(category)
			this.titleNewCategory = ""
		})
	}

	openDeleteCategory(identifier:string) {
		this.deleteCategoryIdentifier = identifier
		this.openAlert(true, "Êtes-vous sur de vouloir supprimer cette catégorie ?", () => {
			this.displayConfirm = false
			const index = this.book.categories.findIndex(category => category.identifier === identifier)

			if (index > -1) {
				this.deleteCategory(this.book.categories[index])
			}
		}, () => {
			this.deleteCategoryIdentifier = ""
			this.displayConfirm = false
		})
	}

	deleteCategory(category:Category) {
		this.book.deleteCategory(session.user.getToken(), category.identifier, () => {
			for (let i = 0; i < this.nodes.length; i++) {
				this.nodes[i].removeCategory(category.identifier)
			}
			this.generateDisplayNode()
		})
	}

	createNode() {
		const node = new Node(this.book.identifier, "New node", "")

		if (this.currentCategoryIdentifier !== "") {
			node.addCategory(this.currentCategoryIdentifier)
		}

		node.record(session.user.getToken(), () => {
			this.nodes.push(node)
			this.generateDisplayNode()
			this.editNode(node)
		})
	}

	saveCurrentNode() {
		this.currentNode.update(session.user.getToken(), () => {
			console.log("save node ok")
			this.generateDisplayNode()
			this.setEditNodeMode(false) /* TODO necessary to do a feedback */
		})
	}

	openDeleteCurrentNode(identifier:string) {
		this.deleteNodeIdentifier = identifier
		this.openAlert(true, "Êtes-vous sur de vouloir supprimer ce noeud ?", () => {
			this.displayConfirm = false
			this.deleteCurrentNode()
		}, () => {
			this.deleteNodeIdentifier = ""
			this.displayConfirm = false
		})
	}

	deleteCurrentNode() {
		this.currentNode.delete(session.user.getToken(), () => {
			const index = this.nodes.findIndex(node => node.identifier === this.deleteNodeIdentifier)

			if (index > -1) {
				this.nodes.splice(index, 1)
			}
			this.currentNode = new Node("", "", "")
			this.deleteNodeIdentifier = ""
			this.generateDisplayNode()
		})
	}
}
