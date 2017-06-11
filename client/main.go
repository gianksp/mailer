/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package main

import (
	"log"
	"fmt"
	"time"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "github.com/gianksp/mailer/proto"
)

 func Submit(c pb.MailingServiceClient) {
	 content := &pb.Content{Type:"plain/text",Value:"Your App is ready!"}
	 from    := &pb.Email{Name:"Jarvis",Address:"jarvis@glofox.com"}
	 to := make([]*pb.Email,0)
	 to  = append(to, &pb.Email{Name:"Gian",Address:"gianksp@gmail.com"})

	 envelope := &pb.Envelope{Subject:"App Ready for sale!",Content:content,From:from,To:to}
	 r, err := c.Send(context.Background(), envelope)
	 if err != nil {
		 log.Fatalf("could not greet: %v %v", err, r)

	 }
	 fmt.Printf("%+v\n", r)
 }

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(os.Getenv("SRV")+":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMailingServiceClient(conn)

	ticker := time.NewTicker(time.Millisecond * 500)

	for t := range ticker.C {
			Submit(c)
	    fmt.Println("Tick at", t)
	}

  fmt.Println("Ticker stopped")

}
