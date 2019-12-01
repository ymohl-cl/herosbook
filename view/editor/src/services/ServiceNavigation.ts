import router from "@/router"

/**
 * Service to manage navigation between views and rewrite url
 */
class Navigation {
	public changeView(route:string):void{
		router.push(route)
	}

	public replaceView(route:string):void{
		router.replace(route)
	}

	public isSameRoute(route:string):boolean {
		return router.currentRoute.name === route
	}
}

const navigation = new Navigation()

export default navigation
