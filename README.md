This is a simple command-line interface program that generates a random list (of user specified length) of Go language features/associated tools. The goal is to provide the user with a concrete set of features/tools that they must incorporate into their next project. This will hopefully help users learn and/or better their of understanding of Go and how to use these features and tools more effectively.

Enter the following into your terminal at the project root directory:

To add a feature:
./bin/add -cat {specified category} -ft {new feature name}

To remove a feature:
./bin/remove -cat {specified category} -ft {specify feature}

To generate a list of required features for your next miniproj:
./bin/gen -nf {enter number of features e.g. 3}
