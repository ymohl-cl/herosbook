<template>
	<v-alert :value="showAlert" style="z-index: 99" absolute :type="current.type" transition="scale-transition">
		{{ current.message }}
	</v-alert>
</template>
<script src="./AlertMessage.ts"></script>
