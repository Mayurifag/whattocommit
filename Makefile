ARGS = $(filter-out $@,$(MAKECMDGOALS))

%:
	@:

include ./.dockerdev/*.mk
