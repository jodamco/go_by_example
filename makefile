new_example:
	mkdir $(name) && touch $(name)/$(name).go && code $(name)/$(name).go

push_example:
	git add . && git commit -m $(m)