import Vue from "vue"
import Component from "vue-class-component"

export class ItemMenu {
	title: string
	f: () => void

	constructor(title: string, f: () => void) {
		this.title = title
		this.f = f
	}
}

function buildMenu(items: any[], disconnectF: () => void): ItemMenu[] {
	const menu: ItemMenu[] = []

	items.forEach((item:ItemMenu) => {
		menu.push(new ItemMenu(item.title, item.f))
	})
	menu.push(new ItemMenu("Disconnect", disconnectF))
	return menu
}

@Component({
	props: {
		connected: { type: Boolean, default: false, required: true },
		pseudo: { type: String, default: "", required: true },
		buildMenu: { type: Array, required: true },
		disconnectFunction: { type: Function, required: true },
		loginFunction: { type: Function, required: true },
		registerFunction: { type: Function, required: true },
	},
})
export default class BaseHeader extends Vue {
	menu: ItemMenu[] = []

	disconnect() {
		this.$props.disconnectFunction()
	}

	register() {
		this.$props.registerFunction()
	}

	login() {
		this.menu = buildMenu(this.$props.buildMenu, this.$props.disconnectFunction)
		this.$props.loginFunction()
	}
}
