import Vue from "vue"
import Component from "vue-class-component"
import BaseHeader from "@/components/BaseHeader/BaseHeader.vue"
import { ItemMenu } from "@/components/BaseHeader/BaseHeader.ts"
import AlertMessage, * as Alert from "@/components/AlertMessage/AlertMessage.ts"
import AlertVue from "@/components/AlertMessage/AlertMessage.vue"
import navigation from "@/services/ServiceNavigation"
import session, { Session } from "@/services/ServiceSession"
import * as route from "@/router"

@Component({
	name: "App",
	components: { BaseHeader, AlertVue },
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
		this.session.alert.addAlert("test", Alert.ErrorMessage)
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
