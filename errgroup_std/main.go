package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	in := []Input{
		{Data: 1},
		{Data: 2},
		{Data: 3},
		{Data: 4},
		{Data: 5},
	}

	res, err := handler(context.Background(), in)
	printResult(res, err)
	fmt.Println("-----------------")
	inBad := []Input{
		{Data: 1},
		{Data: 0},
		{Data: 3},
		{Data: 4},
	}
	res, err = handler(context.Background(), inBad)
	printResult(res, err)
}

func printResult(res []Res, err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	for i, r := range res {
		fmt.Printf("%d. result: %s\n", i, r.Message)
	}
}

func handler(ctx context.Context, in []Input) ([]Res, error) {  
    results := make([]Res, len(in))  
    g, _ := errgroup.WithContext(ctx)  
   
    for i, data := range in {  
        g.Go(func() error {
			i, d := i, data.Data // NB to copy i and data.Data to goroutine scope
			fmt.Printf("processing data: %d\n", d)
            result, err := handle(d)
            if err != nil {  
                return err
            }  
            results[i] = result

            return nil  
        })  
    }  
   
    if err := g.Wait(); err != nil {
        return nil, err  
    }  
    return results, nil  
}

func handle(data int) (Res, error) { // it can be context-aware but we shouldn't call ctx.Done
	if data <= 0 {
		return Res{}, fmt.Errorf("data must be greater than 0")
	}
	return Res{Message: fmt.Sprintf("success: %d", data)}, nil
}

type Input struct {
	Data int
}

type Res struct {
	Message string
}