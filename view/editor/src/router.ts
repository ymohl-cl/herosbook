import Vue from 'vue';
import Router from 'vue-router';
import Home from '@/views/home/home.vue';

Vue.use(Router);

export const landingPagePath = '/landing'
export const resumePagePath = '/'
export const loginPagePath = '/login'
export const registerPagePath = '/register'

export default new Router({
	mode: 'history',
	base: process.env.BASE_URL,
	routes: [
/*		{
			path: '/',
			name: 'home',
			component: Home,
		},*/
		{
			path: resumePagePath,
			name: 'resume-home-page',
			component: () => import('@/views/ResumeHome/ResumeHome.vue'),
		},
		{
			path: loginPagePath,
			name: 'login-view',
			component: () => import('@/views/ViewLogin/ViewLogin.vue'),
		},
		{
			path: registerPagePath,
			name: 'register-view',
			component: () => import('@/views/ViewRegister/ViewRegister.vue'),
		},
		{
			path: '/book/:id',
			name: 'book',
			component: () => import('@/views/book/book.vue'),
		},
		{
			path: landingPagePath,
			name: 'landing-page',
			component: () => import('@/views/LandingPage/LandingPage.vue'),
		},
		{ path: '*', redirect: '/' },
	],
});
