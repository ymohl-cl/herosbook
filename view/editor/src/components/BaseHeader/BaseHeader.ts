import Vue from 'vue';
import Component from 'vue-class-component';
import navigation from '@/services/ServiceNavigation'
import session, {Session} from '@/services/ServiceSession';
import * as route from '@/router'


export class ItemMenu {
	title: string
	f: () => void

	constructor(title: string, f: () => void) {
		this.title = title
		this.f = f
	}
}

@Component
export default class BaseHeader extends Vue {
	session: Session = session
	menu: ItemMenu[] = [
		new ItemMenu("profil", this.loadProfile),
		new ItemMenu("disconnect", this.session.user.disconnect)]

	loadProfile() {
		console.log("load profile view")
	}
	register() {
		navigation.replaceView(route.registerPagePath)
	}
	login() {
		navigation.replaceView(route.loginPagePath)
	}
}