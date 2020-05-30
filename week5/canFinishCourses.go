package main;


type Course struct{
	N int
	PreReqs []int
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	courses := make([]*Course, numCourses);
	for i, course := range courses{
		course = &Course { N: i, PreReqs: []int {}};
		courses[i] = course;
	}

	for _, row := range prerequisites{
		preReq := row[0];
		to := row[1];

		course := courses[to];
		course.PreReqs = append(course.PreReqs, preReq);
	}

	for _, course := range courses{
		visitedNodes := make([]bool, len(courses));
		if cycleDetection(courses, course, visitedNodes){
			return false;
		}
	}

	return true;		
}

func cycleDetection(courses []*Course , course *Course, visitedNodes []bool) bool{
	if len(course.PreReqs) == 0{	
		return false;
	}

	if visitedNodes[course.N]{
		return true;
	}
	visitedNodes[course.N] = true;

	
	for true {
		preReq:= course.PreReqs[0];
		if visitedNodes[preReq]{
			return true;
		}

		if cycleDetection(courses, courses[preReq], visitedNodes){
			return true;
		}

		course.PreReqs = course.PreReqs[1:];
		if len(course.PreReqs) == 0{
			break;
		}
	}

	for i, _ := range visitedNodes{
		visitedNodes[i] = false;
	}

	return false;
}