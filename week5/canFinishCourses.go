package main;

type Course struct{
	N int
	PreReqs []int
}

func canFinish(numCourses int, prerequisites [][]int) bool {

	courses := make([]*Course, numCourses);
	for _, row := range prerequisites{
		preReq := row[0];
		to := row[1];

		course := courses[to];
		if course == nil{
			course = &Course {N: to, PreReqs: []int { preReq }};
			courses[to] = course;
		}else{
			course.PreReqs = append(course.PreReqs, preReq);
		}
	}
	for i, course := range courses{
		if course == nil{
			course = &Course { N: i, PreReqs: []int {}};
			courses[i] = course;
		}
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

	if course != nil && visitedNodes[course.N]{
		return true;
	}
	visitedNodes[course.N] = true;

	for _, preReq := range course.PreReqs{
		if visitedNodes[preReq]{
			return true;
		}

		if cycleDetection(courses, courses[preReq], visitedNodes){
			return true;
		}
	}

	for i, _ := range visitedNodes{
		visitedNodes[i] = false;
	}

	return false;
}

func main(){
	result := canFinish(4, [][]int {{0,1},{3,1},{1,3},{3,2}});
	// result := canFinish(2, [][]int {{0,1}, {1,0}});
	print(result);
}