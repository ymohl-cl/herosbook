import Vue from "vue"
import Router from "vue-router"

Vue.use(Router)

export const landingPagePath = "/landing"
export const resumePagePath = "/"
export const loginPagePath = "/login"
export const registerPagePath = "/register"
export const bookPagePath = "/book/:id"
export const createBookPagePath = "/book/create"

export function buildBookPagePath(identifier: string): string {
	return `/book/${identifier}`
}

export default new Router({
	mode: "history",
	base: process.env.BASE_URL,
	routes: [
		{
			path: resumePagePath,
			name: "resume-home-page",
			component: () => import("@/views/ResumeHome/ResumeHome.vue"),
		},
		{
			path: loginPagePath,
			name: "login-view",
			component: () => import("@/views/ViewLogin/ViewLogin.vue"),
		},
		{
			path: registerPagePath,
			name: "register-view",
			component: () => import("@/views/ViewRegister/ViewRegister.vue"),
		},
		{
			path: bookPagePath,
			name: "book-view",
			component: () => import("@/views/ViewBook/ViewBook.vue"),
		},
		{
			path: createBookPagePath,
			name: "book-create",
			component: () => import("@/views/CreateBook/CreateBook.vue"),
		},
		{
			path: landingPagePath,
			name: "landing-page",
			component: () => import("@/views/LandingPage/LandingPage.vue"),
		},
		{ path: "*", redirect: "/" },
	],
})
