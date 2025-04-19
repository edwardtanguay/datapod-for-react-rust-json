import { action, Action, thunk, Thunk } from "easy-peasy";
import { FrontendFlashcard } from "../../types";
import * as dataModel from "../dataModel";
import { StoreModel } from "../store";

export interface FlashcardModel {
	// state
	frontendFlashcards: FrontendFlashcard[];
	numberOfAnswersShown: number;

	// actions
	setFrontendFlashcards: Action<this, FrontendFlashcard[]>;
	saveFrontendFlashcard: Action<this, FrontendFlashcard>;
	incrementAnswersShown: Action<this>;

	// thunks
	loadFlashcardsThunk: Thunk<this>;
	toggleFrontendFlashcardThunk: Thunk<
		this,
		FrontendFlashcard,
		void,
		StoreModel
	>;
}

export const flashcardModel: FlashcardModel = {
	// state
	frontendFlashcards: [],
	numberOfAnswersShown: 0,

	// actions
	setFrontendFlashcards: action((state, flashcards) => {
		state.frontendFlashcards = structuredClone(flashcards);
	}),
	saveFrontendFlashcard: action((state, frontendFlashcard) => {
		const index = state.frontendFlashcards.findIndex(
			(s) => s.suuid === frontendFlashcard.suuid
		);
		if (index !== -1) {
			state.frontendFlashcards[index] =
				structuredClone(frontendFlashcard);
		}
	}),
	incrementAnswersShown: action((state) => {
		state.numberOfAnswersShown++;
	}),

	// thunks
	loadFlashcardsThunk: thunk((actions) => {
		(async () => {
			const _frontendFlashcards = dataModel.getFlashcards();
			actions.setFrontendFlashcards(_frontendFlashcards);
		})();
	}),
	toggleFrontendFlashcardThunk: thunk(
		(actions, frontendFlashcard, helpers) => {
			frontendFlashcard.isOpen = !frontendFlashcard.isOpen;
			if (frontendFlashcard.isOpen) {
				actions.incrementAnswersShown();
				helpers
					.getStoreActions()
					.mainModel.setMessage(
						`Number of times an answer was shown: ${
							helpers.getState().numberOfAnswersShown
						}`
					);
			}
			actions.saveFrontendFlashcard(frontendFlashcard);
		}
	),
};
