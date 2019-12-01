<template>
	<v-container>
		<div class="mt-12">Todo : Change by a true editor</div>
		<v-textarea solo :value="text" @change="change"></v-textarea>
	</v-container>
</template>
<script src="./TextEditor.ts"></script>
