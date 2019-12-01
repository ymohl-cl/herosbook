import Vue from "vue"
import Component from "vue-class-component"
import BaseHeader from "@/components/BaseHeader/BaseHeader.vue"
import { ItemMenu } from "@/components/BaseHeader/BaseHeader.ts"
import AlertMessage from "@/components/AlertMessage/AlertMessage.vue"
import navigation from "@/services/ServiceNavigation"
import session, { Session } from "@/services/ServiceSession"
import * as route from "@/router"

@Component({
	name: "App",
	components: { BaseHeader, AlertMessage },
})
export default class App extends Vue {
	session:Session = session

	mounted() {
		// TODO: implement ping method to start or not the application
		if (!session.user.isConnected()) {
			navigation.replaceView(route.landingPagePath)
		} else {
			navigation.replaceView(route.resumePagePath)
		}
		this.session.alert = this.$refs.AlertMessage as AlertMessage
	}

	menuBuilder(): ItemMenu[] {
		return [
			new ItemMenu("Home", () => {
				navigation.replaceView(route.resumePagePath)
			}),
			new ItemMenu("Profile", () => {
				console.log("load profile view")
			}),
		]
	}

	loginView(): void {
		navigation.replaceView(route.loginPagePath)
	}

	registerView(): void {
		navigation.replaceView(route.registerPagePath)
	}

	disconnect(): void {
		session.user.disconnect()
		navigation.replaceView(route.landingPagePath)
	}
}
