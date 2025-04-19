import { z } from "zod";

export const FlashcardSchema = z.object({
	suuid: z
		.string()
		.trim()
		.length(6, "suuid must be exactly 6 characters long")
		.regex(
			/^[A-Za-z0-9]+$/,
			"suuid can only contain uppercase/lowercase letters and numbers"
		)
		.min(1, "suuid cannot be empty"),
	category: z.string().trim().min(1, "Category cannot be empty"),
	front: z.string().trim().min(1, "Front cannot be empty"),
	back: z.string().trim().min(1, "Back cannot be empty"),
});

export const FrontendFlashcardSchema = FlashcardSchema.extend({
	isOpen: z.boolean(),
});

export type Flashcard = z.infer<typeof FlashcardSchema>;
export type FrontendFlashcard = z.infer<typeof FrontendFlashcardSchema>;
