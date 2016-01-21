
build:
	cp ../build/Linux/entrykit .
	docker build -t entrykit-example .
	rm entrykit

run:
	docker run -p 8000:80 --name entrykit-example entrykit-example

extra:
	cat extra.conf | docker exec -i entrykit-example sh -c "cat > /demo/includes/extra.conf"
