import { FlashcardSchema, FrontendFlashcard } from "../types";
import rawFlashcards from "../../parseddata/flashcards.json";

export const getFlashcards = () => {
	const frontendFlashcards: FrontendFlashcard[] = [];
	for (const rawFlashcard of rawFlashcards) {
		const parseResult = FlashcardSchema.safeParse(rawFlashcard);
		if (parseResult.success) {
			const { suuid, category, front, back } = parseResult.data;
			const frontendFlashcard: FrontendFlashcard = {
				suuid: suuid.trim(),
				category: category.trim(),
				front: front.trim(),
				back: back.trim(),
				isOpen: false,
			};
			frontendFlashcards.push(frontendFlashcard);
		} else {
			let r = "";
			r += `INVALID FLASHCARD IN IMPORT: ${JSON.stringify(
				rawFlashcard,
				null,
				2
			)}\n`;
			parseResult.error.errors.forEach((err) => {
				r += `Error in field "${err.path.join(".")}" - ${
					err.message
				}\n`;
			});
			console.error(r);
		}
	}
	return frontendFlashcards;
};
